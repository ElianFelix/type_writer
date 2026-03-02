// Helper functions to help process text ie game text setup


export function processInputText(text) {
  const output = text.split('').map(
    (l, idx) => {
      return { letter: l, status: idx == 0 ? 'cursor' : '' }
    }
  )
  return output
}
