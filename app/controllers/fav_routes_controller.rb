class FavRoutesController < ApplicationController
  before_action :set_fav_route, only: [:show, :update, :destroy]

  # GET /fav_routes
  def index
    @fav_routes = FavRoute.all

    render json: @fav_routes
  end

  # GET /fav_routes/1
  def show
    render json: @fav_route
  end

  # POST /fav_routes
  def create
    @fav_route = FavRoute.new(fav_route_params)

    if @fav_route.save
      render json: @fav_route, status: :created, location: @fav_route
    else
      render json: @fav_route.errors, status: :unprocessable_entity
    end
  end

  # PATCH/PUT /fav_routes/1
  def update
    if @fav_route.update(fav_route_params)
      render json: @fav_route
    else
      render json: @fav_route.errors, status: :unprocessable_entity
    end
  end

  # DELETE /fav_routes/1
  def destroy
    @fav_route.destroy
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_fav_route
      @fav_route = FavRoute.find(params[:id])
    end

    # Only allow a trusted parameter "white list" through.
    def fav_route_params
      params.require(:fav_route).permit(:user_id, :polyline1, :polyline2, :polyline3, :polyline4, :polyline5, :count)
    end
end
