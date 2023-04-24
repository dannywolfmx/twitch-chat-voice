
const synth = window.speechSynthesis

export const readMessage = (message: string) => {
    console.log("Talking")
    const speech = new SpeechSynthesisUtterance(message)

    speech.rate = 1.2
    synth.speak(speech)
}
