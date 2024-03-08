import { Navigator } from '../App'
import './Admin.css'

type AdminProps = {
  navigate: Navigator
}

function Admin(props: AdminProps) {

  return (
    <div>
      <h3 class="_show-title">Admin Page</h3>
      <div class="_buttons" >
      <button class="_button" onClick={()=>props.navigate("allslides")}> Slides</button>
      <button class="_button" onClick={()=>props.navigate("slideshow")}> SlideShow</button>
      <button class="_button" onClick={()=>props.navigate("slideshows")}> SlideShows</button>
      <button class="_button" onClick={()=>props.navigate("newslideshow")}> New SlideShow</button>
      <button class="_button" onClick={()=>props.navigate("newslide")}> New Slide</button>
    </div>
    </div>
  )
}

export default Admin
