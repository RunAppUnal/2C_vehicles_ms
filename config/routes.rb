Rails.application.routes.draw do
  resources :fav_routes do
    collection do
      get :my_favRoutes
    end
  end

  resources :vehicles do
  	collection do
  		get :my_vehicles
  		get :find_vehicle
  		get :count_my_vehicles
  		get :total_vehicles
  	end
  end
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
