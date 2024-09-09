export default function Notify(displayView: Function) {
  displayView(true)
  setTimeout(() => {
    displayView(false)
  }, 2500)
}

