'use strict';

function _interopDefault (ex) { return (ex && (typeof ex === 'object') && 'default' in ex) ? ex['default'] : ex; }

var backendCommon = require('@backstage/backend-common');
var config = require('@backstage/config');
var knex2 = _interopDefault(require('knex'));
var pluginAuthBackend = require('@backstage/plugin-auth-backend');
var pluginCatalogBackend = require('@backstage/plugin-catalog-backend');
var pluginIdentityBackend = require('@backstage/plugin-identity-backend');
var pluginScaffolderBackend = require('@backstage/plugin-scaffolder-backend');
var rest = require('@octokit/rest');
var Docker = _interopDefault(require('dockerode'));
var pluginProxyBackend = require('@backstage/plugin-proxy-backend');
var pluginTechdocsBackend = require('@backstage/plugin-techdocs-backend');

async function createPlugin({
  logger,
  database,
  config
}) {
  return await pluginAuthBackend.createRouter({logger, config, database});
}

async function createPlugin$1({
  logger,
  database
}) {
  const locationReader = new pluginCatalogBackend.LocationReaders(logger);
  const db = await pluginCatalogBackend.DatabaseManager.createDatabase(database, {logger});
  const entitiesCatalog = new pluginCatalogBackend.DatabaseEntitiesCatalog(db);
  const locationsCatalog = new pluginCatalogBackend.DatabaseLocationsCatalog(db);
  const higherOrderOperation = new pluginCatalogBackend.HigherOrderOperations(entitiesCatalog, locationsCatalog, locationReader, logger);
  backendCommon.useHotCleanup(module, pluginCatalogBackend.runPeriodically(() => higherOrderOperation.refreshAllLocations(), 1e4));
  return await pluginCatalogBackend.createRouter({
    entitiesCatalog,
    locationsCatalog,
    higherOrderOperation,
    logger
  });
}

async function createPlugin$2({logger}) {
  return await pluginIdentityBackend.createRouter({logger});
}

async function createPlugin$3({logger}) {
  const cookiecutterTemplater = new pluginScaffolderBackend.CookieCutter();
  const craTemplater = new pluginScaffolderBackend.CreateReactAppTemplater();
  const templaters = new pluginScaffolderBackend.Templaters();
  templaters.register("cookiecutter", cookiecutterTemplater);
  templaters.register("cra", craTemplater);
  const filePreparer = new pluginScaffolderBackend.FilePreparer();
  const githubPreparer = new pluginScaffolderBackend.GithubPreparer();
  const preparers = new pluginScaffolderBackend.Preparers();
  preparers.register("file", filePreparer);
  preparers.register("github", githubPreparer);
  const githubClient = new rest.Octokit({auth: process.env.GITHUB_ACCESS_TOKEN});
  const publisher = new pluginScaffolderBackend.GithubPublisher({client: githubClient});
  const dockerClient = new Docker();
  return await pluginScaffolderBackend.createRouter({
    preparers,
    templaters,
    publisher,
    logger,
    dockerClient
  });
}

async function createPlugin$4({
  logger,
  config
}) {
  return await pluginProxyBackend.createRouter({logger, config});
}

async function createPlugin$5({
  logger,
  config
}) {
  const generators = new pluginTechdocsBackend.Generators();
  const techdocsGenerator = new pluginTechdocsBackend.TechdocsGenerator();
  generators.register("techdocs", techdocsGenerator);
  const directoryPreparer = new pluginTechdocsBackend.DirectoryPreparer();
  const preparers = new pluginTechdocsBackend.Preparers();
  preparers.register("dir", directoryPreparer);
  const publisher = new pluginTechdocsBackend.LocalPublish();
  const dockerClient = new Docker();
  return await pluginTechdocsBackend.createRouter({
    preparers,
    generators,
    publisher,
    dockerClient,
    logger,
    config
  });
}

var _a;
function makeCreateEnv(loadedConfigs) {
  const config2 = config.ConfigReader.fromConfigs(loadedConfigs);
  return (plugin) => {
    const logger = backendCommon.getRootLogger().child({type: "plugin", plugin});
    const knexConfig = {
      client: "sqlite3",
      connection: ":memory:",
      useNullAsDefault: true
    };
    const database = knex2(knexConfig);
    database.client.pool.on("createSuccess", (_eventId, resource) => {
      resource.run("PRAGMA foreign_keys = ON", () => {
      });
    });
    return {logger, database, config: config2};
  };
}
async function main() {
  const configs = await backendCommon.loadBackendConfig();
  const configReader = config.ConfigReader.fromConfigs(configs);
  const createEnv = makeCreateEnv(configs);
  const catalogEnv = backendCommon.useHotMemoize(module, () => createEnv("catalog"));
  const scaffolderEnv = backendCommon.useHotMemoize(module, () => createEnv("scaffolder"));
  const authEnv = backendCommon.useHotMemoize(module, () => createEnv("auth"));
  const identityEnv = backendCommon.useHotMemoize(module, () => createEnv("identity"));
  const proxyEnv = backendCommon.useHotMemoize(module, () => createEnv("proxy"));
  const techdocsEnv = backendCommon.useHotMemoize(module, () => createEnv("techdocs"));
  const service = backendCommon.createServiceBuilder(module).loadConfig(configReader).addRouter("/catalog", await createPlugin$1(catalogEnv)).addRouter("/scaffolder", await createPlugin$3(scaffolderEnv)).addRouter("/auth", await createPlugin(authEnv)).addRouter("/identity", await createPlugin$2(identityEnv)).addRouter("/techdocs", await createPlugin$5(techdocsEnv)).addRouter("/proxy", await createPlugin$4(proxyEnv));
  await service.start().catch((err) => {
    console.log(err);
    process.exit(1);
  });
}
(_a = module.hot) == null ? void 0 : _a.accept();
main().catch((error) => {
  console.error(`Backend failed to start up, ${error}`);
  process.exit(1);
});
