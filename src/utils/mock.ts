import type {Account} from '../types/accounts';

export function loadMockAccount(path: string): Account[] {
    const fs = require('fs');
    const rawdata = fs.readFileSync(path);
    const accounts: Account[] = JSON.parse(rawdata);

    return accounts;
}


async function sleepUpToSecond() {
    const min = 0;
    const max = 1000;
    const randMS = Math.random() * (max - min) + min;
    return new Promise((r) => setTimeout(r, randMS));
}

export async function* fetchAccountResponseGenerator(accounts: Account[]) {
    for (let i = 0; i < accounts.length; i++) {
        await sleepUpToSecond();
        yield accounts[i];
    }
}