import { createSignal, onCleanup, onMount } from 'solid-js'
import './SlideShow.css'
import SlideButton, { ButtonDirection } from './SlideButton'
import { useSocket } from '../hooks/useSocket'

function SlideShow() {
  const [imageSrc, setImageSrc] = createSignal<string>("")
  const getCat = async () => {
    const req = await fetch("https://api.thecatapi.com/v1/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1",
    {
      method: "GET",
      redirect: "follow"
    });
    const catSon = await req.json()
    const url = catSon[0]["url"]
    setImageSrc(url)
  }
  const handleMessage = (message: any) => {
    console.log(message);
  }
  const [ws, sendMessage] = useSocket(handleMessage);
  const handleRightArrow = () => sendMessage("right");
  const handleLeftArrow = () => sendMessage("left");
  
  onMount(()=> {
    getCat();
  });
  onCleanup(()=> {
    ws.close();
  });

  return (
    <div class="_slideshow">
    <img class="_slide" src={imageSrc()} height="400" width="600"/>
    <div class="_slide-button-div">
    <SlideButton direction={ButtonDirection.LEFT} action={handleLeftArrow} />
    <SlideButton direction={ButtonDirection.RIGHT} action={handleRightArrow} />
    </div>
    </div>
  )
}

export default SlideShow
