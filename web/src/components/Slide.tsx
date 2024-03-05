import { createResource, createSignal,  onMount } from 'solid-js'
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
const getSlide = async (_url: string) => {
    const resp = await fetch(_url,{
      method: "GET",
      redirect: "follow",})
    const blob = await resp.blob()
    return URL.createObjectURL(blob)
    }
  
function Slide(props: SlideProps) {
  const url = `http://127.0.0.1:8000/blob/${props.SlideData.slide.url}`;
  const [imageSrc, setImageSrc] = createSignal<string>()
  setImageSrc(url)
  const [slide] = createResource(imageSrc, getSlide)

  return (
    <div class="_slide-div">
      {slide.loading ? 
     <div class="slide-placeholder">Loading...</div> :
    <img class="_slide" src={slide()} height="400" width="600"/> }
    </div>
  )
}

export default Slide
