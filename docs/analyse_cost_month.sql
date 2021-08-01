
DROP PROCEDURE IF EXISTS analyse_cost_month;
DELIMITER //
-- description 支出统计
-- author lujiang
-- created_at 2021-08-01
create procedure analyse_cost_month(in ym int)
begin
    DECLARE start_date date;
    DECLARE end_date date;


    set start_date = date_format(concat(ym,'01'),'%Y-%m-%d');
    set end_date = last_day(start_date);

    SET SQL_SAFE_UPDATES = 0;

    -- 按【user_id, ym, category_id】 进行统计
    delete from wealth_cost_statistics_months where `year_month` = ym;

    insert into wealth_cost_statistics_months(`year`, user_id,`year_month`,category_id,total)
    select date_format(occur_date,'%Y'),user_id,date_format(occur_date,'%Y%m') ym, category_id, sum(amount) total
    from wealth_cost_details
    where occur_date between start_date and end_date
    group by date_format(occur_date,'%Y'),user_id,date_format(occur_date,'%Y%m'), category_id;

end
//
DELIMITER ;
