import { createSignal, onCleanup, onMount } from 'solid-js'
import './SlideShow.css'
import SlideButton, { ButtonDirection } from '../components/SlideButton'
import Slide, { SlideData } from '../components/Slide'
import { useSocket } from '../hooks/useSocket'

export type SlideShowData = {
  slideshow_data: {
    id: number,
    title: string,
    created_at: string,
  },
  slides: SlideData[]
}

type ShowState = {
  slideshow: number,
  slide: number
}

enum DIRECT {
  LEFT = -1,
  RIGHT = 1
 }

function SlideShow() {
  const [slideShowData, setSlideShowData] = createSignal<SlideShowData>();
  const [slideShowPosition, setSlideShowPosition] = createSignal<number>(0);
  const [slideShowState, setSlideShowState] = createSignal<ShowState|null>(null)
  const getSlideShowData = async () => {
    const res = await fetch("http://localhost:8000/api/slideshow",
    {
      method: "GET",
      redirect: "follow"
    });
    const data = await res.json()
    console.log(data)
    setSlideShowData(data);
  }
  onMount(()=>{
    getSlideShowData()
  })
  const handleMessage = (message: string) => {
    try {
    const msgJson = JSON.parse(message)
    setSlideShowState(msgJson["state"])
    } catch {
      console.log("error reading socket traffic")
    }
  }
  const [ws, sendMessage] = useSocket(handleMessage);
  const directSlide = (state: ShowState | null, direct: DIRECT) : string => {
    const _default: ShowState = {slideshow: 1, slide: 0}
    if (state == null) {
      return JSON.stringify(_default);
    }
    if (direct == DIRECT.LEFT && state.slide > 0) {
      state.slide--;
      return JSON.stringify(state);
    }
    if (direct == DIRECT.RIGHT) {
      state.slide++;
      return JSON.stringify(state);
    }
    return JSON.stringify(_default);
  }
  const handleRightArrow = () => sendMessage(directSlide(slideShowState(), DIRECT.RIGHT));
  const handleLeftArrow = () => sendMessage(directSlide(slideShowState(), DIRECT.LEFT));
  const slides = () => slideShowData()?.slides.map(slide => <Slide slide={slide}/>)
  const getCurrentSlide = () => { 
    const allSlides = slides();
    const idx = slideShowState()?.slide ? slideShowState()?.slide as number : 0 as number
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
  const slideNumber = () => slideShowState() ? slideShowState()?.slide : "#";

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
    <h3>{slideNumber()}</h3>
    <SlideButton direction={ButtonDirection.RIGHT} action={handleRightArrow} />
    </div>
    </div>
    </div>
  )
}

export default SlideShow
