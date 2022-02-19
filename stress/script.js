import http from 'k6/http';
import { sleep } from 'k6';


// max rate is 100 jokes per minute
// 2 users = 50 jokes each per minute
// avg 5 jokes / rq = 10 rq / min
// => sleep minimum 5s between requests

export default function () {
    const url = 'http://localhost:8080/joke'
    // around 5s between request
    // randomness helps desynchronize users
    const sleepTime = Math.floor(5 + Math.random() * 2)
    const payload = JSON.stringify({
        // max jokes nb is 10, expect some errors
        qty: Math.floor(Math.random() * 11),
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.request('GET', url, payload, params);

    sleep(sleepTime);
}
