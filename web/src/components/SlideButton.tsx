import "./SlideButton.css"

type SlideButtonProps = {
  direction: boolean;
}

function SlideButton(props: SlideButtonProps) {
  const direction = props.direction ? 'left' : 'right';

  return (
    <div class="_slide-button">
    <p>
      {direction}
      </p> 
    </div>
  )
}

export default SlideButton
