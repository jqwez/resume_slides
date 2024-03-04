import { JSXElement, createSignal, onCleanup, onMount } from 'solid-js'
import './SlideShow.css'
import SlideButton, { ButtonDirection } from './SlideButton'
import Slide, { SlideType } from './Slide'
import { useSocket } from '../hooks/useSocket'

type SlideShowData = {
  id: number,
  title: string,
  created_at: string,
  slides: SlideType[]
}

function SlideShow() {
  const [slideShowData, setSlideShowData] = createSignal<SlideShowData>()
  const getSlideShowData = async () => {
    const res = await fetch("http://localhost:8000/slideshow/default",
    {
      method: "GET",
      redirect: "follow"
    });
    const data = await res.json()
    setSlideShowData(data);
    console.log(data)
    console.log(data.slides)
  }
  onMount(()=>{
    getSlideShowData()
  })
  const handleMessage = (message: any) => {
    console.log(message);
  }
  const [ws, sendMessage] = useSocket(handleMessage);
  const handleRightArrow = () => sendMessage("right");
  const handleLeftArrow = () => sendMessage("left");
  const slides: JSXElement[] = [1, 2].map(num => <Slide blobURL={"http://localhost:8000/blob/cat.jpg"} num={num} />)

  const showTitle = () => slideShowData() ? slideShowData()?.title : "Title Holder";

  onCleanup(()=> {
    ws.close();
  });

  return (
    <div>
      <h3 class="_show-title">{showTitle()}</h3>
    <div class="_slideshow">
      {slides[0]}
    <div class="_slide-button-div">
    <SlideButton direction={ButtonDirection.LEFT} action={handleLeftArrow} />
    <SlideButton direction={ButtonDirection.RIGHT} action={handleRightArrow} />
    </div>
    </div>
    </div>
  )
}

export default SlideShow
