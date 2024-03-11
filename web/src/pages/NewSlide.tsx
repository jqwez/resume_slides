import { Navigator } from '../App'
import AdminNav from '../components/AdminNav'
import { useEnvironmentVariable } from '../hooks/useEnvironment'
import './NewSlide.css'

type NewSlideProps = {
  navigate: Navigator
}
function NewSlide(props: NewSlideProps) {
  const newSlideRequest = async(show: string, fileT:string, file: File) => {
    const formData = new FormData();
    formData.append('file', file)
    formData.append('filet', fileT)
    formData.append('show', show)
    const baseUrl = useEnvironmentVariable("container_ip");
    const url = baseUrl ? `https://${baseUrl}/api/admin/slide/new` : `http://localhost:8000/api/admin/slide/new`
    fetch(url, {
      method: "POST",
      body: formData
    })
  } 
  const handleSubmit = (e: MouseEvent) => {
    e.preventDefault();
    const titleInput = document.getElementById("slide-title") as HTMLInputElement;
    const slideTitle = titleInput.value;
    const fileTitleInput = document.getElementById("slide-title") as HTMLInputElement;
    const fileTitle = fileTitleInput.value;
    const fileInput = document.getElementById("filePicker") as HTMLInputElement;
    const file = fileInput?.files?.[0]
    if (file) {
      newSlideRequest(slideTitle, fileTitle, file);
    } else {
    }
  } 
  
  return (
    <div>
      <AdminNav navigate={props.navigate} />
      <h3 class="_show-title">New Slide</h3>
      <div id="_new-slide-form">
      <label for="slide-title">Title</label><input id="slide-title" name="title" />
      <label for="filePicker">First Slide</label>
      <input type="file" name="file" id="filePicker" accept=".jpg, .jpeg, .png" />
      <button onClick={handleSubmit} id="submit-button">Upload</button>
      </div>
    </div>
  )
}

export default NewSlide