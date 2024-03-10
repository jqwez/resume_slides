export function useStoredPage() {
   function getPage() {
    return localStorage.getItem("stored-page");
  }
  
  function setPage(page: string) {
    localStorage.setItem("stored-page", page);
}
 
  return {getStoredPage: getPage, setStoredPage: setPage}
}