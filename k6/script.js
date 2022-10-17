import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 },
    { duration: '1m30s', target: 40 },
    { duration: '5m00s', target: 50 },
    { duration: '4m00s', target: 25 },
    { duration: '10m00s', target: 100 },
    { duration: '40s', target: 0 },
  ],
};

export default function () {
  const res = http.get('https://app.onboarding-counter-api.us-east-1.general.stag.wildlife.io/api/v1');
  check(res, { 'status was 200': (r) => r.status == 200 });
}