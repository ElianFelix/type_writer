import axios from 'axios'

const API_BASE_URL = 'http://localhost:1323/'


const appApi =
  axios.create({
    baseURL: API_BASE_URL,
  })

//
// Auth
//
export async function login(loginInfo = {}) {
  try {
    const response = await appApi.post('/login', loginInfo)
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}
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

export async function updateUser(updUser = {}, authToken) {
  try {
    const response = await appApi.put(`/users/${updUser.id}`,
                                      updUser,
                                      { headers: { Authorization: authToken }})
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function deleteUser(id = 0, authToken) {
  try {
    const response = await appApi.delete(`/users/${id}`,
                                         { headers: { Authorization: authToken }})
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

export async function createActivity(activity = {}, authToken) {
  try {
    const response = await appApi.post('/activities',
                                       activity,
                                       { headers: { Authorization: authToken }})
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function updateActivity(updActivity = {}, authToken) {
  try {
    const response = await appApi.put(`/activities/${updActivity.id}`,
                                      updActivity,
                                      { headers: { Authorization: authToken }})
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function deleteActivity(id = 0, authToken) {
  try {
    const response = await appApi.delete(`/activities/${id}`,
                                               { headers: { Authorization: authToken }})
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

export async function createText(text = {}, authToken) {
  try {
    const response = await appApi.post('/texts',
                                       text,
                                       { headers: { Authorization: authToken }})
    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function updateText(updText = {}, authToken) {
  try {
    const response = await appApi.put(`/texts/${updText.id}`,
                                      updText,
                                      { headers: { Authorization: authToken }})
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

export async function createScore(score = {}, authToken) {
  try {
    const response = await appApi.post('/scores',
                                       score,
                                       { headers: { Authorization: authToken }})

    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function updateScore(updScore = {}, authToken) {
  try {
    const response = await appApi.put(`/scores/${updScore.id}`,
                                      updScore,
                                      { headers: { Authorization: authToken }})

    return response.data
  } catch(error) {
    console.log(error)
  }
  return null
}

export async function deleteScore(id = 0, authToken) {
  try {
    const response = await appApi.delete(`/scores/${id}`, { headers: { Authorization: authToken }})

    return response.data
  } catch(error) {
    console.log(error)
  }
  return false
}

