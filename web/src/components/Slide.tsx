import { createSignal,  onMount } from 'solid-js'
import './Slide.css'

export type SlideData = {
  slide: {
    id: number,
    title: string,
    url: string,
    slideshow_id: number,
    created_at: string,
  },
  position: number,
}

type SlideProps = {
  SlideData: SlideData;
}

function Slide(props: SlideProps) {
  const url = `http://127.0.0.1:8000/blob/${props.SlideData.slide.url}`;
  const [imageSrc, setImageSrc] = createSignal<string>("")
  const getCat = async () => {
    const req = await fetch(url,
    {
      method: "GET",
      redirect: "follow",
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
