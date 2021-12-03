import http from 'k6/http';
import { sleep, check } from 'k6';
import { Counter } from 'k6/metrics';

export const requests = new Counter('http_reqs');

export const options = {
    stages: [
      { target: 20, duration: '1m' },
      { target: 15, duration: '1m' },
      { target: 0, duration: '1m' },
    ],
    thresholds: {
      requests: ['count < 100'],
    },
  };
  

export default function () {
  const url = 'http://localhost:8010/users';
  
   const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  let res = http.get(url);

  check(res, {
    "is status 200": (r) => r.status === 200   
  });

  sleep(Math.random()* 5)
}