import './SlideShows.css'
import { createSignal, onMount } from 'solid-js'
import { SlideData } from '../components/Slide'
import { Navigator } from '../App'
import Thumbnail from '../components/Thumbnail'
import AdminNav from '../components/AdminNav'
import { useEnvironmentVariable } from '../hooks/useEnvironment'


type SlideShowEditorProps = {
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

function SlideShowEditor(props: SlideShowEditorProps) {
  const [slideShowData, setSlideShowData] = createSignal<SlideShowData>();
  const getSlideShowData = async () => {
    const baseUrl = useEnvironmentVariable("container_ip", "http://127.0.0.1:8000")
    const res = await fetch(`${baseUrl}/admin/slideshows`,
    {
      method: "GET",
      redirect: "follow"
    });
    const data = await res.json()
    setSlideShowData(data);
    console.log(slideShowData());
  }
  onMount(()=>{
    getSlideShowData()
  })

  const bunch = [0, 1, 2, 3, 4, 5, 7, 21, 1, 1, 3].map(_=><Thumbnail imageUrl='cat.jpg'/>)
  return (
    <> 
    <AdminNav navigate={props.navigate} />
    <h3>Slide Show Editor</h3>
       <div class="thumbnail-grid">
      {bunch}
    </div>
</>

  )
}

export default SlideShowEditor
