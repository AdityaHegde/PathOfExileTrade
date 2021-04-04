import * as React from "react";
import {useEffect, useState} from "react";
import {fetchGet, fetchTry} from "@typescript/utils/fetch";

export function Dashboard() {
  const [sample, setSample] = useState<any>({})

  useEffect(() => {
    fetchTry(fetchGet("/api/sample")).then(sampleResp => setSample(sampleResp));
  }, []);

  return (
    <>
      <div>User: {sample.UserName}</div>
      <div>Sample: {sample.Sample}</div>
    </>
  )
}
