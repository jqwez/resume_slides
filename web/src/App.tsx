import './App.css'
import SlideShow from "./pages/SlideShow"
import Admin from "./pages/Admin";
import NewSlideShow from './pages/NewSlideShow';
import SlideShows from './pages/SlideShows';
import { JSXElement, createSignal, onMount } from 'solid-js';
import { useStoredPage } from './hooks/useStoredPage';


export type Navigator = (dest: string) => void;

function App() {
  const {getPage, setPage} = useStoredPage();
  const [currentPage, setCurrentPage] = createSignal<JSXElement>(<SlideShow navigate={navigate} />);
  const pages: Record<string, JSXElement> = {
    "slideshow": <SlideShow navigate={navigate} />,
    "slideshowadmin": <Admin navigate={navigate} />,
    "slideshows": <SlideShows navigate={navigate} />,
    "newslideshow": <NewSlideShow navigate={navigate} />,
  }
  
  function navigate(dest: string) {
    const destination = pages[dest];
    if (destination == undefined) {
      console.log("Not a destination");
      return
    }
    setPage(dest);
    setCurrentPage(pages[dest]);
  }
  onMount(()=> {
    const page = getPage();
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
