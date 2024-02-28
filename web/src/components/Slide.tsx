import { createSignal,  onMount } from 'solid-js'
import './SlideShow.css'


type SlideProps = {
  num: any;
}

function Slide(props: SlideProps) {
  const [imageSrc, setImageSrc] = createSignal<string>("")
  const getCat = async () => {
    const req = await fetch("http://localhost:8000/blob",
    {
      method: "GET",
      redirect: "follow"
    });
    const blob = await req.blob()
    setImageSrc(URL.createObjectURL(blob))
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
