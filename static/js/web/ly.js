class Home{
    inits(){
        this.acrid()
    }
    // 热门推荐随机颜色
    acrid(){
        const colors = ["#FE2D46", "#F60000", "#FAA90E"]
        let acids = $('.acrid').find(".acrid_label")
        for (let i = 0; i < 5; i++) {
            if (i < 3) {
                $(acids[i]).css({"backgroundColor": colors[i]})
            }
        }
    }
}

let home = new Home
home.inits()