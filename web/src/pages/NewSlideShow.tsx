import { Navigator } from '../App'
import AdminNav from '../components/AdminNav'
import { useEnvironmentVariable } from '../hooks/useEnvironment'
import './NewSlideShow.css'

type NewSlideShowProps = {
  navigate: Navigator
}
function NewSlideShow(props: NewSlideShowProps) {
  const newSlideShowRequest = async(show: string, fileT:string, file: File) => {
    const formData = new FormData();
    formData.append('file', file)
    formData.append('filet', fileT)
    formData.append('show', show)
    const baseUrl = useEnvironmentVariable("container_ip", "http://127.0.0.1:8000")
    fetch(`${baseUrl}/newslideshow`, {
      method: "POST",
      body: formData
    })
  } 
  const handleSubmit = (e: MouseEvent) => {
    e.preventDefault();
    const titleInput = document.getElementById("slideshow-title") as HTMLInputElement;
    const showTitle = titleInput.value;
    const fileTitleInput = document.getElementById("slide-title") as HTMLInputElement;
    const fileTitle = fileTitleInput.value;
    const fileInput = document.getElementById("filePicker") as HTMLInputElement;
    const file = fileInput?.files?.[0]
    if (file) {
      newSlideShowRequest(showTitle, fileTitle, file);
    } else {
    }
  } 
  
  return (
    <div>
      <AdminNav navigate={props.navigate} />
      <h3 class="_show-title">New SlideShow</h3>
      <div id="_new-slideshow-form">
      <label for="slideshow-title">Title</label><input id="slideshow-title" name="title" />
      <label for="filePicker">First Slide</label>
      <input id="slide-title" type="text" name="slide-title"></input>
      <input type="file" name="file" id="filePicker" accept=".jpg, .jpeg, .png" />
      <button onClick={handleSubmit} id="submit-button">Upload</button>
      </div>
    </div>
  )
}

export default NewSlideShow