import { Navigator } from "../App";
import "./AdminNav.css"

type AdminNavProps = {
  navigate: Navigator
}

function AdminNav(props: AdminNavProps) {
  const pages = [
    ["slideshows", "SlideShows"],
    ["slideshowadmin", "Admin"],
    ["newslideshow", "New SlideShow"],
    ["newslide", "New Slide"],
    ["allslides", "All Slidse"],
  ] 
  const buttons = pages.map(page=><button class="admin-nav-button" onClick={()=>props.navigate(page[0])}>{page[1]}</button>)
  return (
    <div class="">
      {buttons}
    </div>
  )
}

export default AdminNav