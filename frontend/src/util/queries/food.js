import axios from "axios";

const url = import.meta.env.VITE_API_URL;
const token = localStorage.getItem("token");

export async function inventory() {
  const endpoint = `${url}/api/food?type=all`;

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

  return response;
}

export async function getById(id) {
  const endpoint = `${url}/api/food?id=${id}`;

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

export async function updateInventory(id, value) {
  const endpoint = `${url}/api/inventory?id=${id}&value=${value}`;

  console.log(id, value);

  const response = await fetch(endpoint, {
    method: "GET",
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

export async function addItem(item) {
  const formData = new FormData();
  const endpoint = `${url}/api/food`;

  console.log(item);
  formData.append("name", item.name);
  formData.append("desc", item.desc);
  formData.append("price", item.price);
  formData.append("stock", item.stock);

  const response = await fetch(endpoint, {
    method: "POST",
    headers: {
      Authorization:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA0MTEzOTksInVzZXJuYW1lIjoia2hhc21vZGFuIn0.jP9zbmAo9JFa0is112bMiyz_l22QuC42TWuOISSOkEQ",
    },
    "Content-Type": "application/x-www-form-urlencoded; boundary=&",

    body: formData,
  })
    .then((res) => {
      return res.json();
    })
    .catch((err) => {
      return err;
    });

  return response;
}

export async function updateItem(item) {
  const formData = new FormData();
  const endpoint = `${url}/api/food`;

  formData.append("id", item.id);
  formData.append("name", item.name);
  formData.append("desc", item.desc);
  formData.append("price", item.price);
  formData.append("stock", item.stock);

  const response = await fetch(endpoint, {
    method: "PATCH",
    headers: {
      Authorization: token,
    },
    "Content-Type": "application/x-www-form-urlencoded; boundary=&",

    body: formData,
  })
    .then((res) => {
      return res.json();
    })
    .catch((err) => {
      return err;
    });

  return response;
}
