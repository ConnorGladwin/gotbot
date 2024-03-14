const url = import.meta.env.VITE_API_URL;
const token = localStorage.getItem("token");
const userId = localStorage.getItem("id");

export async function signUp(user) {
  const formData = new FormData();

  formData.append("username", user.username);
  formData.append("firstName", user.firstName);
  formData.append("lastName", user.lastName);
  formData.append("email", user.email);
  formData.append("password", user.password);

  const endpoint = `${url}/auth/sign-up`;

  const response = fetch(endpoint, {
    method: "POST",
    body: formData,
    "Content-Type": "multipart/form-data",
  })
    .then((res) => res.json())
    .then((json) => {
      return json;
    })
    .catch((err) => {
      return "error:" + err;
    });

  return response;
}

export async function login(user) {
  const formData = new FormData();

  formData.append("id", user.id);
  formData.append("password", user.password);

  const endpoint = `${url}/auth/login`;

  const response = fetch(endpoint, {
    method: "POST",
    body: formData,
    "Content-Type": "multipart/form-data",
  })
    .then((res) => res.json())
    .then((json) => {
      return json;
    })
    .catch((err) => {
      return "error:" + err;
    });

  return response;
}

export async function validateRoute(token) {
  const endpoint = `${url}/auth/validate?token=${token}`;

  const response = fetch(endpoint, {
    method: "GET",
  })
    .then((res) => {
      return res.status;
    })
    .catch((err) => {
      return false;
    });

  return response;
}

export async function getUser() {
  const endpoint = `${url}/api/user?id=${userId}`;

  const response = await fetch(endpoint, {
    method: "GET",
    headers: {
      Authorization: token,
      "Content-Type": "text/plain",
    },
  })
    .then((res) => {
      return res.json();
    })
    .catch((err) => {
      return err;
    });

  return response[0];
}

export async function updateUser(user) {
  const endpoint = `${url}/api/user?id=${userId}&firstName=${user.firstName}&lastName=${user.lastName}&username=${user.username}`;

  const response = await fetch(endpoint, {
    method: "PATCH",
    headers: {
      Authorization: token,
    },
    "Content-Type": "application/json",
  })
    .then((res) => {
      return res.json();
    })
    .catch((err) => {
      return err;
    });

  return response;
}
