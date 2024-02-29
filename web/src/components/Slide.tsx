import { createSignal,  onMount } from 'solid-js'
import './SlideShow.css'


export type SlideType = {
  id: number,
  title: string,
  url: string,
  slideshow_id: number,
  position: number,
  created_at: string,
}

type SlideProps = {
  blobURL: string,
  num: any;
}

function Slide(props: SlideProps) {
  const [imageSrc, setImageSrc] = createSignal<string>("")
  const getCat = async () => {
    const req = await fetch(props.blobURL,
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
    <div class="_slide-div">
    <img class="_slide" src={imageSrc()} height="400" width="600"/>
    </div>
  )
}

export default Slide
