import { createSignal,  onMount } from 'solid-js'
import './SlideShow.css'


type SlideProps = {
  num: any;
}

function Slide(props: SlideProps) {
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
  
  onMount(()=> {
    getCat();
  });

  return (
    <div class="_slide-dive">
    <img class="_slide" src={imageSrc()} height="400" width="600"/>
    </div>
  )
}

export default Slide
