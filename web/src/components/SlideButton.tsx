import "./SlideButton.css"

export enum ButtonDirection {
  LEFT,
  RIGHT
}

type SlideButtonProps = {
  direction: ButtonDirection;
  action: () => void;
}

function SlideButton(props: SlideButtonProps) {
  const direction = props.direction == ButtonDirection.LEFT ? '🢀' : '🢂';
  return (
    <div class="_slide-button" onClick={props.action}>
    <p>
      {direction}
      </p> 
    </div>
  )
}

export default SlideButton