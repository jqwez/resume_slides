import { JSXElement, onCleanup } from 'solid-js'
import './SlideShow.css'
import SlideButton, { ButtonDirection } from './SlideButton'
import Slide from './Slide'
import { useSocket } from '../hooks/useSocket'

function SlideShow() {
  const handleMessage = (message: any) => {
    console.log(message);
  }
  const [ws, sendMessage] = useSocket(handleMessage);
  const handleRightArrow = () => sendMessage("right");
  const handleLeftArrow = () => sendMessage("left");
  const slides: JSXElement[] = [1, 2].map(num => <Slide num={num} />)
  onCleanup(()=> {
    ws.close();
  });

  return (
    <div class="_slideshow">
      {slides[0]}
    <div class="_slide-button-div">
    <SlideButton direction={ButtonDirection.LEFT} action={handleLeftArrow} />
    <SlideButton direction={ButtonDirection.RIGHT} action={handleRightArrow} />
    </div>
    </div>
  )
}

export default SlideShow
