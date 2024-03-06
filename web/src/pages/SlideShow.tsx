import { createSignal, onCleanup, onMount } from 'solid-js'
import './SlideShow.css'
import SlideButton, { ButtonDirection } from '../components/SlideButton'
import Slide, { SlideData } from '../components/Slide'
import { useSocket } from '../hooks/useSocket'
import { Navigator } from '../App'


type SlideShowProps = {
  navigate: Navigator
}

export type SlideShowData = {
  slideshow_data: {
    id: number,
    title: string,
    created_at: string,
  },
  slides: SlideData[]
}

function SlideShow(props: SlideShowProps) {
  const [slideShowData, setSlideShowData] = createSignal<SlideShowData>();
  const [slideShowPosition, setSlideShowPosition] = createSignal<number>(0);
  const getSlideShowData = async () => {
    const res = await fetch("http://localhost:8000/slideshow/1",
    {
      method: "GET",
      redirect: "follow"
    });
    const data = await res.json()
    setSlideShowData(data);
  }
  onMount(()=>{
    getSlideShowData()
  })
  const handleMessage = (message: any) => {
    console.log(message);
  }
  const [ws, sendMessage] = useSocket(handleMessage);
  const handleRightArrow = () => props.navigate("slideshowadmin");
  const handleLeftArrow = () => sendMessage("left");
  const slides = () => slideShowData()?.slides.map(
    slide => <Slide 
      SlideData={slide}
      />)
  const getCurrentSlide = () => { 
    const allSlides = slides();
    const idx = slideShowPosition();
    if (allSlides == undefined) return undefined;
    if (allSlides.length > idx) {
      return allSlides[idx];
    } else {
      setSlideShowPosition(allSlides.length - 1);
      if (slideShowPosition() >= 0) {
        return allSlides[idx];
      }
    }
  } 

  const showTitle = () => slideShowData() ? slideShowData()?.slideshow_data?.title : "Title Holder";

  onCleanup(()=> {
    ws.close();
  });

  return (
    <div>
      <h3 class="_show-title">{showTitle()}</h3>
    <div class="_slideshow">
      {getCurrentSlide()}
    <div class="_slide-button-div">
    <SlideButton direction={ButtonDirection.LEFT} action={handleLeftArrow} />
    <SlideButton direction={ButtonDirection.RIGHT} action={handleRightArrow} />
    </div>
    </div>
    </div>
  )
}

export default SlideShow
