import './SlideShows.css'
import { Navigator } from '../App'
const Thumbnail = lazy(()=> import('../components/Thumbnail'))
import AdminNav from "../components/AdminNav"
import { lazy } from 'solid-js'

type SlideShowsProps = {
  navigate: Navigator
}

function SlideShows(props: SlideShowsProps) {
  const bunch = [0, 1, 2, 3, 4, 5, 7, 21, 1, 1, 3].map((_)=><Thumbnail imageUrl='cat.jpg'/>)
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
