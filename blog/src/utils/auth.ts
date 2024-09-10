type Profile = {
  isLogged: boolean
  username: string
  email: string
}

export default async function FetchUserData() : Promise<Profile> {

  const endpoint = import.meta.env.PUBLIC_API_HOST + "/auth/blog/profile"
  const user = { isLogged: false, username: "", email: "" }

  try {
    const response = await fetch(endpoint, {
      method: "GET",
      credentials: "include",
    })

    const result = await response.json()

    if (result?.status == "success") {
      return {
        isLogged: true,
        username: result?.data?.username,
        email: result?.data?.email
      }
    }
    
    return user
  } catch(err) {
    return user
  }

}
