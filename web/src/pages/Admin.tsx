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
      <button class="_button" onClick={()=>props.navigate("slideshow")}> SlideShow</button>
      <button class="_button" onClick={()=>props.navigate("slideshows")}> SlideShows</button>
    </div>
    </div>
  )
}

export default Admin
