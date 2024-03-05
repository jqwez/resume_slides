import { Navigator } from "../App";
import "./AdminNav.css"
import { onMount, createSignal } from "solid-js";

type AdminNavProps = {
  navigate: Navigator
}

function AdminNav(props: AdminNavProps) {
  const pages = [
    ["slideshows", "SlideShows"],
    ["slideshowadmin", "Admin"],
    ["newslideshow", "New SlideShow"]
  ] 
  const buttons = pages.map(page=><button class="admin-nav-button" onClick={()=>props.navigate(page[0])}>{page[1]}</button>)
  return (
    <div class="">
      {buttons}
    </div>
  )
}

export default AdminNav