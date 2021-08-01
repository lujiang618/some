
DROP PROCEDURE IF EXISTS analyse_cost_week;
DELIMITER //
-- description 支出统计
-- author lujiang
-- created_at 2021-08-01
create procedure analyse_cost_week(in call_date date)
begin
    DECLARE start_date date;
    DECLARE end_date date;

    set start_date = DATE_SUB(call_date,INTERVAL WEEKDAY(call_date) DAY);
    set end_date = DATE_SUB(call_date,INTERVAL WEEKDAY(call_date)-6 DAY);

    SET SQL_SAFE_UPDATES = 0;

    -- 按【user_id, y, star_date,end_date, category_id】 进行统计
    delete from wealth_cost_statistics_weeks where start_date = start_date;
    
    insert into wealth_cost_statistics_weeks(`year`, start_date, end_date, user_id, category_id, total)
    select 
        date_format(occur_date,'%Y'),
        start_date, 
        end_date, 
        user_id,
        category_id,
        sum(amount) total
    from wealth_cost_details
    where occur_date between start_date and end_date
    group by date_format(occur_date,'%Y'),user_id,category_id;

end
//
DELIMITER ;
