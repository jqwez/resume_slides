import { createSignal, onMount } from 'solid-js'
import { Navigator } from '../App'
import AdminNav from '../components/AdminNav'
import { useEnvironmentVariable } from '../hooks/useEnvironment'

type AllSlidesProps = {
  navigate: Navigator
}
function AllSlides(props: AllSlidesProps) {
  const getAllSlides = async () => {
    const baseUrl = useEnvironmentVariable("container_ip", "http://127.0.0.1:8000")
    const res = await fetch(`${baseUrl}/api/admin/slide/all`,
    {
      method: "GET",
      redirect: "follow"
    });
    const data = await res.json()
    console.log(data)
    setSlides(data.slides)
  }
  const [slides, setSlides] = createSignal()
  onMount(()=>getAllSlides())
  console.log(slides())
  const SlideElements = () => {
    const _slides = slides() as Array<Object>;
    if (_slides != null) {
      return _slides.map(_ => <SlideRow/>)
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