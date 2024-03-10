import { createEffect, createSignal, onCleanup } from 'solid-js';

export function useListenForAdmin(cb: ()=>void) {
  const [sequence, setSequence] = createSignal<string[]>([]);
  const secretSequence = ["KeyA", "KeyD", "KeyM", "KeyI", "KeyN"];
    createEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      setSequence((prevSequence) => [...prevSequence, event.code]);
      if (sequence().toString() === secretSequence.toString()) {
        cb()
      }
    };
    window.addEventListener("keydown", handleKeyDown);
    onCleanup(() => {
      window.removeEventListener("keydown", handleKeyDown);
    });
  }, [sequence, secretSequence]);
}