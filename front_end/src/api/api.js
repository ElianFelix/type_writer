import axios from 'axios'

const API_BASE_URL = 'http://localhost:1323/'


const appApi =
  axios.create({
    baseURL: API_BASE_URL,
  })


//
// Users
//
export async function getUsers() {
  try {
    const response = await appApi.get('/users')
    return response.data.users
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function getUserByIdOrUsername(keyword = 0) {
  try {
    const response = await appApi.get(`/users/${keyword}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function createUser(user = {}) {
  try {
    const response = await appApi.post('/users', user)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function updateUser(updUser = {}) {
  try {
    const response = await appApi.put(`/users/${updUser.id}`, updUser)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function deleteUser(id = 0) {
  try {
    const response = await appApi.delete(`/users/${id}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return false
}

//
// Activities
//
export async function getActivities() {
  try {
    const response = await appApi.get('/activities')
    return response.data.activities
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function getActivityByIdOrName(keyword = 0) {
  try {
    const response = await appApi.get(`/activities/${keyword}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function createActivity(activity = {}) {
  try {
    const response = await appApi.post('/activities', activity)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function updateActivity(updActivity = {}) {
  try {
    const response = await appApi.put(`/activities/${updActivity.id}`, updActivity)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function deleteActivity(id = 0) {
  try {
    const response = await appApi.delete(`/activities/${id}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return false
}

//
// Texts
//
export async function getTexts() {
  try {
    const response = await appApi.get('/texts')
    return response.data.texts
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function getTextByIdOrTitle(keyword = 0) {
  try {
    const response = await appApi.get(`/texts/${keyword}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function createText(text = {}) {
  try {
    const response = await appApi.post('/texts', text)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function updateText(updText = {}) {
  try {
    const response = await appApi.put(`/texts/${updText.id}`, updText)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function deleteText(id = 0) {
  try {
    const response = await appApi.delete(`/texts/${id}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return false
}

//
// Activity Results / Scores
//
export async function getScores() {
  try {
    const response = await appApi.get('/scores')
    return response.data.scores
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function getScoreById(id = 0) {
  try {
    const response = await appApi.get(`/scores/${id}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function createScore(score = {}) {
  try {
    const response = await appApi.post('/scores', score)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function updateScore(updScore = {}) {
  try {
    const response = await appApi.put(`/scores/${updScore.id}`, updScore)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function deleteScore(id = 0) {
  try {
    const response = await appApi.delete(`/scores/${id}`)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return false
}

