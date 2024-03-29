import { useEnvironmentVariable } from "./useEnvironment";

export function useSocket(handler: (message: any)=>void): [WebSocket, (message: any) => void] {
  const baseUrl = useEnvironmentVariable("container_ip")
  const url = baseUrl ? `wss://${baseUrl}/ws` : `ws://localhost:8000/ws`
  const ws = new WebSocket(url)
  ws.onopen = () => {
    //ws.send("client connected");
    console.log("connected to server");
  }
  ws.onmessage = (event) => {
    handler(event.data);
  }
  ws.onclose = () => {
    //ws.send("Client Disconnecting");
    console.log("Disconnected from server");
  }
  const sendMessage = (message: any) => {
    if (ws.readyState === WebSocket.OPEN) {
      ws.send(message);
    } else {
      console.log("WebSocket is not connected");
    }
  }


  return [ws, sendMessage];
}