require 'test_helper'

class FavRoutesControllerTest < ActionDispatch::IntegrationTest
  setup do
    @fav_route = fav_routes(:one)
  end

  test "should get index" do
    get fav_routes_url, as: :json
    assert_response :success
  end

  test "should create fav_route" do
    assert_difference('FavRoute.count') do
      post fav_routes_url, params: { fav_route: { count: @fav_route.count, polyline1: @fav_route.polyline1, polyline2: @fav_route.polyline2, polyline3: @fav_route.polyline3, polyline4: @fav_route.polyline4, polyline5: @fav_route.polyline5, user_id: @fav_route.user_id } }, as: :json
    end

    assert_response 201
  end

  test "should show fav_route" do
    get fav_route_url(@fav_route), as: :json
    assert_response :success
  end

  test "should update fav_route" do
    patch fav_route_url(@fav_route), params: { fav_route: { count: @fav_route.count, polyline1: @fav_route.polyline1, polyline2: @fav_route.polyline2, polyline3: @fav_route.polyline3, polyline4: @fav_route.polyline4, polyline5: @fav_route.polyline5, user_id: @fav_route.user_id } }, as: :json
    assert_response 200
  end

  test "should destroy fav_route" do
    assert_difference('FavRoute.count', -1) do
      delete fav_route_url(@fav_route), as: :json
    end

    assert_response 204
  end
end
