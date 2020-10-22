import express from 'express';
import { Logger } from 'winston';

interface RouterOptions {
    logger: Logger;
}
declare function createRouter(options: RouterOptions): Promise<express.Router>;

export { RouterOptions, createRouter };
