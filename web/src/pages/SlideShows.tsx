import './SlideShows.css'
import { createSignal, onMount } from 'solid-js'
import { SlideData } from '../components/Slide'
import { Navigator } from '../App'
import Thumbnail from '../components/Thumbnail'
import AdminNav from "../components/AdminNav"


type SlideShowsProps = {
  navigate: Navigator
}

type SlideShowData = {
  slideshow_data: {
    id: number,
    title: string,
    created_at: string,
  },
  slides: SlideData[]
}

function SlideShows(props: SlideShowsProps) {
  /*const [slideShowData, setSlideShowData] = createSignal<SlideShowData>();
  const getSlideShowData = async () => {
    const res = await fetch("http://localhost:8000/admin/slideshows",
    {
      method: "GET",
      redirect: "follow"
    });
    const data = await res.json()
    setSlideShowData(data);
  }
  onMount(()=>{
    getSlideShowData()
  })*/

  const bunch = [0, 1, 2, 3, 4, 5, 7, 21, 1, 1, 3].map(el=><Thumbnail imageUrl='cat.jpg'/>)
  return (
    <>
    <div>
    <AdminNav navigate={props.navigate}/>
    </div>
    <div class="thumbnail-grid">
      {bunch}
    </div>
</>
  )
}

export default SlideShows
