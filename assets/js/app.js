angular.module('app', ['ngRoute']);
require('./routes');
require('./apiInterceptor');

require('../controllers/rootController');
require('../controllers/loginController');

require('../services/authTokenService');
require('../services/sessionService');
