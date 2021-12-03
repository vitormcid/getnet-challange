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
  const url = 'http://localhost:8010/user';
  const payload = JSON.stringify({
    name: 'user teste',
    email: 'user@test.com',
    password: 'Test!A1gkl',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  let res = http.post(url, payload, params);

  check(res, {
    "is status 201": (r) => r.status === 201    
  });

  sleep(Math.random()* 5)
}