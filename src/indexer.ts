'use strict';

import type {Account} from './types/accounts';
import {loadMockAccount, fetchAccountResponseGenerator} from './utils/mock';

// var args = process.argv;
startIndexer("./input.json");


function startIndexer(inputPath: string) {
    const accounts = loadMockAccount(inputPath);

    fetchAccountResponseGenerator(accounts);

    console.log(accounts);
}

