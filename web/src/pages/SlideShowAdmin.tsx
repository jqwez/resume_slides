import { Navigator } from '../App'
import './SlideShowAdmin.css'

type SlideShowAdminProps = {
  navigate: Navigator
}
function SlideShowAdmin(props: SlideShowAdminProps) {
  return (
    <div>
      <h3 class="_show-title">SlideShow Admin Page</h3>
      <div class="_buttons" >
      <button class="_button" onClick={()=>props.navigate("slideshoweditor")}> Upload Slide</button>
      <button class="_button" onClick={()=>props.navigate("slideshows")}> SlideShows</button>
    </div>
    </div>
  )
}

export default SlideShowAdmin
