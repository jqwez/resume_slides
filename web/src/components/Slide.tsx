import { Suspense, createResource, createSignal } from 'solid-js'
import './Slide.css'
import { useEnvironmentVariable } from '../hooks/useEnvironment';

export type SlideData = {
    id: number,
    title: string,
    url: string,
    created_at: string,
    position: number,
}

type SlideProps = {
  slide: SlideData;
}
const getSlide = async (_url: string) => {
    const resp = await fetch(_url,{
      method: "GET",
      redirect: "follow",})
    const blob = await resp.blob()
    return URL.createObjectURL(blob)
    }
  
function Slide(props: SlideProps) {
  const baseUrl = useEnvironmentVariable("container_ip")
  const url = baseUrl ? `https://${baseUrl}/api/blob/${props.slide.url}` : `http://localhost:8000/api/blob/${props.slide.url}`;
  const [imageSrc, setImageSrc] = createSignal<string>()
  setImageSrc(url)
  const [slide] = createResource(imageSrc, getSlide)

  return (
    <div class="_slide-div">
    <Suspense fallback={<div class="slide-placeholder">Loading...</div>}>
    <img class="_slide" src={slide()} height="400" width="600"/> 
    </Suspense>
    </div>
  )
}

export default Slide
