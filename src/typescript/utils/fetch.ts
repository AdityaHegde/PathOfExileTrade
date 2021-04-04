export function fetchGet(url) {
  return fetch(url, {
    method: "GET",
    headers: {},
  });
}

export function fetchPost(url, bodyJson) {
  return fetch(url, {
    method: "POST",
    headers: {},
    body: JSON.stringify(bodyJson),
  });
}

export async function fetchTry(fetchPromise: Promise<Response>): Promise<any> {
  try {
    const resp = await fetchPromise;
    if (resp.ok) {
      return await resp.json();
    } else {
      console.log(resp.statusText);
    }
  } catch (err) {
    console.log(err);
  }

  return null;
}
