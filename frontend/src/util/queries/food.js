import axios from "axios";

const url = import.meta.env.VITE_API_URL;

export async function inventory() {
  const endpoint = `${url}/api/food?type=all`;

  const response = await fetch(endpoint, {
    method: "GET",
    // Authorization:
    //   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAyMzM2OTIsInVzZXJuYW1lIjoiY29ubm9yeW9ya2dsYWR3aW5AZ21haWwuY29tIn0.cPUpYrv4IzzQFOI28IpuyfryNVyl-8rd5QXIZr5ZJm4",

    headers: {
      Authorization:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA0MTEzOTksInVzZXJuYW1lIjoia2hhc21vZGFuIn0.jP9zbmAo9JFa0is112bMiyz_l22QuC42TWuOISSOkEQ",
      "Content-Type": "text/plain",
    },
  })
    .then((res) => {
      return res.json();
    })
    .catch((err) => {
      return err;
    });

  return response;
}
