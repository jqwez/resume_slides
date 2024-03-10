import "./Thumbnail.css"
import { onMount, createSignal } from "solid-js";

type ThumbnailProps = {
  imageUrl: string
}

function Thumbnail(props: ThumbnailProps) {
  console.log(props);
  const url = `http://127.0.0.1:8000/api/blob/cat.jpg`;
  const [imageSrc, setImageSrc] = createSignal<string>("")
  const getCat = async () => {
    const req = await fetch(url,
    {
      method: "GET",
      redirect: "follow",
    });
    const blob = await req.blob()
    setImageSrc(URL.createObjectURL(blob))
  }
  onMount(()=> {
    getCat();
  });
  return (
    <div class="thumbnail-div">
    <img class="thumbnail" src={imageSrc()} height="200" width="300"/>
    <h4>Title</h4>
    </div>
  )
}

export default Thumbnail