import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
  location: config.locationType,
  rootURL: config.rootURL,
});

Router.map(function() {
  this.route('jobs', function() {
    this.route('job', { path: '/:job_name' }, function() {
      this.route('task-group', { path: '/:name' });
      this.route('definition');
      this.route('versions');
      this.route('deployments');
    });
  });

  this.route('clients', function() {
    this.route('client', { path: '/:node_id' });
  });

  this.route('servers', function() {
    this.route('server', { path: '/:agent_id' });
  });

  this.route('allocations', function() {
    this.route('allocation', { path: '/:allocation_id' }, function() {
      this.route('task', { path: '/:name' }, function() {
        this.route('logs');
      });
    });
  });

  this.route('settings', function() {
    this.route('tokens');
  });

  if (config.environment === 'development') {
    this.route('freestyle');
  }

  this.route('not-found', { path: '/*' });
});

export default Router;
