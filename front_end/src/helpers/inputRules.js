//
// Centralized file for input validation rules accross application
//
//
export const usernameRules = [
  (v) => /^\w+$/.test(v)
]
export const passwordRules = [
  (v) => /^.{8,}$/.test(v)
]
export const nameRules = [
  (v) => v.length === 0 || /^(\w+|\w+ \w+)$/.test(v)
]
export const emailRules = [
  (v) => /^\w+@\w+\.\w{3}$/.test(v)
]

