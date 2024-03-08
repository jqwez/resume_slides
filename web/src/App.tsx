import './App.css'
import SlideShow from "./pages/SlideShow"
import Admin from "./pages/Admin";
import NewSlideShow from './pages/NewSlideShow';
import SlideShows from './pages/SlideShows';
import { JSXElement, createSignal, onMount } from 'solid-js';
import { useStoredPage } from './hooks/useNavigator';
import { useListenForAdmin } from './hooks/useGoToAdmin';
import NewSlide from './pages/NewSlide';
import AllSlides from './pages/AllSlides';


export type Navigator = (dest: string) => void;

function App() {
  useListenForAdmin(()=>navigate("slideshowadmin"))
  const {getStoredPage, setStoredPage} = useStoredPage();
  const [currentPage, setCurrentPage] = createSignal<JSXElement>(<SlideShow />);
  const pages: Record<string, JSXElement> = {
    "slideshow": <SlideShow />,
    "slideshowadmin": <Admin navigate={navigate} />,
    "slideshows": <SlideShows navigate={navigate} />,
    "newslideshow": <NewSlideShow navigate={navigate} />,
    "newslide": <NewSlide navigate={navigate} />,
    "allslides": <AllSlides navigate={navigate} />,
  }
  
  function navigate(dest: string) {
    const destination = pages[dest];
    if (destination == undefined) {
      console.log("Not a destination");
      return
    }
    setStoredPage(dest);
    setCurrentPage(pages[dest]);
  }
  onMount(()=> {
    const page = getStoredPage();
    console.log(page)
    if (page != undefined) {
      navigate(page)
    };
  })
  setCurrentPage(<Admin navigate={navigate}/>)

  return (
    <>
    {currentPage()}
  </>
  )
}

export default App
