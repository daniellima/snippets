import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 10 },
    { duration: '1m30s', target: 25 },
    { duration: '5m00s', target: 30 },
    { duration: '4m00s', target: 30 },
    { duration: '10m00s', target: 15 },
    { duration: '40s', target: 0 },
  ],
};

export default function () {
  const res = http.post('www.google.com', JSON.stringify({
    count: "asd",
  }))

  check(res, { 'status was 400': (r) => r.status == 400 });
}