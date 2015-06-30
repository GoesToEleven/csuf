$stdout.sync = true

require './setup'

map '/assets' do
  environment = Sprockets::Environment.new
  %w{javascripts stylesheets images templates}.each do |path|
    environment.append_path "./assets/#{path}"
    environment.append_path "./vendor/assets/#{path}"
  end
  run environment
end

use Rack::MethodOverride

map "/" do
  run Html5ImageViewer
end

