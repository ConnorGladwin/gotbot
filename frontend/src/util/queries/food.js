const url = import.meta.env.VITE_API_URL;

export async function inventory() {
  const endpoint = `${url}/auth/validate?type=${all}`;

  const response = fetch(endpoint, {
    method: "GET",
  })
    .then((res) => {
      return res.status;
    })
    .catch((err) => {
      return false;
    });

  console.log(response);

  return response;
}
