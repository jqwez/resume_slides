import { createSignal, onMount } from 'solid-js'
import { Navigator } from '../App'
import AdminNav from '../components/AdminNav'

type AllSlidesProps = {
  navigate: Navigator
}
function AllSlides(props: AllSlidesProps) {
  const getAllSlides = async () => {
    const res = await fetch("http://localhost:8000/api/admin/slide/all",
    {
      method: "GET",
      redirect: "follow"
    });
    const data = await res.json()
    console.log(data)
    setSlides(data.slides)
  }
  const [slides, setSlides] = createSignal(null)
  onMount(()=>getAllSlides())
  console.log(slides())
  const SlideElements = () => {
    if (slides()) {
      return slides().map(<SlideRow/>)
    }
  }
  return (
    <div>
      <AdminNav navigate={props.navigate} />
      {SlideElements()}
      <h3 class="_show-title">All Slide</h3>
    </div>
  )
}

function SlideRow() {
  return "hih"
}

export default AllSlides